use std::io::{BufRead, BufReader, Write};
use std::net::{Shutdown, TcpStream, ToSocketAddrs};
use std::sync::mpsc::{channel, Receiver, Sender};
use std::thread;
use std::time::Duration;

use eframe::egui::Context;

const CONNECT_TIMEOUT: Duration = Duration::from_secs(5);

pub enum NetEvent {
    Connected,
    Line(String),
    Closed(String),
}

pub struct Connection {
    outgoing: Sender<String>,
    incoming: Receiver<NetEvent>,
}

fn connect(addr: &str) -> std::io::Result<(TcpStream, TcpStream)> {
    let mut last_err = None;
    for sock in addr.to_socket_addrs()? {
        match TcpStream::connect_timeout(&sock, CONNECT_TIMEOUT) {
            Ok(stream) => {
                let reader = stream.try_clone()?;
                return Ok((stream, reader));
            }
            Err(e) => last_err = Some(e),
        }
    }
    Err(last_err.unwrap_or_else(|| {
        std::io::Error::new(std::io::ErrorKind::NotFound, "no address found")
    }))
}

impl Connection {
    pub fn open(host: &str, port: &str, ctx: Context) -> Self {
        let addr = format!("{host}:{port}");
        let (event_tx, incoming) = channel();
        let (outgoing, command_rx) = channel::<String>();

        thread::spawn(move || {
            let (mut stream, reader_stream) = match connect(&addr) {
                Ok(pair) => pair,
                Err(e) => {
                    let _ = event_tx.send(NetEvent::Closed(format!("connection failed: {e}")));
                    ctx.request_repaint();
                    return;
                }
            };
            let _ = event_tx.send(NetEvent::Connected);
            ctx.request_repaint();

            let reader_tx = event_tx.clone();
            let reader_ctx = ctx.clone();
            thread::spawn(move || {
                for line in BufReader::new(reader_stream).lines() {
                    let Ok(line) = line else { break };
                    if reader_tx.send(NetEvent::Line(line)).is_err() {
                        return;
                    }
                    reader_ctx.request_repaint();
                }
                let _ = reader_tx.send(NetEvent::Closed("connection closed".into()));
                reader_ctx.request_repaint();
            });

            for command in command_rx {
                if writeln!(stream, "{command}").is_err() {
                    break;
                }
            }
            let _ = stream.shutdown(Shutdown::Both);
        });

        Self { outgoing, incoming }
    }

    pub fn send(&self, command: &str) {
        let clean: String = command.chars().filter(|c| *c != '\n' && *c != '\r').collect();
        let _ = self.outgoing.send(clean);
    }

    pub fn try_recv(&self) -> Option<NetEvent> {
        self.incoming.try_recv().ok()
    }
}

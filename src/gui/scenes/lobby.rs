//pops an internal window in the existing window, w/ address, host, player's username
use eframe::egui::{Ui, Layout, Align};
use crate::gui::components::{button, window, field, row, align};
use crate::state::{AppState, Scene};
use crate::palette::{colored, sized, bold, *};
use crate::network::{Connection, NetEvent};

pub fn connect(ui: &mut Ui, state: &mut AppState) {
    let text_size: f32 = state.settings.text_size;
    let connecting = state.connecting;
    let mut cancel = false;

    window(ui, "Connect to server", &mut state.show_connect, |ui| {
        if let Some(err) = &state.connect_error {
            ui.label(colored(bold(err.as_str()), RED));
        }

        field(ui, &state.settings, sized("Host", text_size), &mut state.host);
        field(ui, &state.settings, sized("Port", text_size), &mut state.port);
        field(ui, &state.settings, sized("Username", text_size), &mut state.username);

        row(ui, |ui| {
            if button(ui, &state.settings, sized("CANCEL", text_size)).clicked() {
                cancel = true;
            }

            align(ui, &Align::Center, |ui| {
                ui.add_enabled_ui(!connecting, |ui| {
                    if button(ui, &state.settings, sized("CONNECT", text_size)).clicked() {
                        state.connect_error = None;
                        state.connection =
                            Some(Connection::open(&state.host, &state.port, ui.ctx().clone()));
                        state.connecting = true;
                        ui.ctx().request_repaint();
                    }
                });
                if connecting {
                    row(ui, |ui| {
                        ui.spinner();
                        ui.label("Connecting...");
                    });
                }
            });
        });
    });

    if cancel || !state.show_connect {
        state.connection = None;
        state.connecting = false;
        state.show_connect = true;
        state.go_back();
    }
}

//pops an internal window in the existing window, w/ address, host, player's username
use eframe::egui::{Ui, Window};
use crate::gui::components::{button, back, window, field};
use crate::state::{AppState, Scene};
use crate::palette::{colored, sized, bold, *};

pub fn  connect(ui: &mut Ui, state: &mut AppState) {
    //the program will reject the name if there is a duplicate returned by the server
    let text_size: f32 = state.settings.text_size;
    let mut cancel: bool = false;

    window(ui, "Connect to server", &mut state.show_connect, |ui| {
        field(ui, &state.settings, sized("Host", text_size), &mut state.host);
        field(ui, &state.settings, sized("Port", text_size), &mut state.port);
        field(ui, &state.settings, sized("Username", text_size), &mut state.stats.name);

        if button(ui, &state.settings, sized("Cancel", text_size)).clicked() {
            cancel = true;
        }
    });

    if cancel || !state.show_connect {
        state.show_connect = true;
        state.go_back();
    }
}

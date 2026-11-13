//pops an internal window in the existing window, w/ address, host, player's username
use eframe::egui::{Ui, Window};
use crate::gui::components::{button, back, window};
use crate::state::{AppState, Scene};
use crate::palette::{colored, sized, bold, *};

pub fn  connect(ui: &mut Ui, state: &mut AppState) {
    //the program will reject the name if there is a duplicate returned by the server
    let text_size: f32 = state.settings.text_size;
    let mut cancel: bool = false;

    window(ui, "Connect to server", &mut state.show_connect, |ui| {
        ui.label(sized(bold("Host"), text_size));
        ui.text_edit_singleline(&mut state.host);
        ui.label(sized(bold("Port"), text_size));
        ui.text_edit_singleline(&mut state.port);
        ui.label(sized(bold("Username"), text_size));
        ui.text_edit_singleline(&mut state.stats.name);

        if ui.button(sized("Cancel", text_size)).clicked() {
            cancel = true;
        }
    });

    if cancel || !state.show_connect {
        state.show_connect = true;
        state.go_back();
    }
}

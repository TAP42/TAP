// byebye. additionally send info to stop server
use eframe::egui::{Ui, ViewportCommand};
use crate::gui::components::centered;
use crate::palette::{colored, sized, bold, *};
use crate::state::AppState;

pub fn  byebye(ui: &mut Ui, state: &mut AppState) {
    let text_size: f32 = state.settings.text_size;

    ui.add_space((ui.available_height() - 80.0) / 2.0);
    centered(ui, |ui| {
        ui.label(colored(sized(bold("Bye bye!"), text_size), YELLOW));
    });
    ui.send_viewport_cmd(ViewportCommand::Close);
}

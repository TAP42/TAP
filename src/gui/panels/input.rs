// footer of the window, with chat + channels [global, room, group] and logs 
use eframe::egui::{Ui, RichText};
use crate::gui::components::{row, chat, label, field, slider};
use crate::state::AppState;
use crate::palette::{colored, sized, bold, *};

pub fn  draw_input(ui: &mut Ui, state: &mut AppState) {
    ui.set_min_height(150.0);
    ui.group(|ui| {
        label(ui, &state.settings, "chat goes here");
    });
}

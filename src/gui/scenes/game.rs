use eframe::egui::{Ui, Panel};
use crate::gui::components::panel;
use crate::gui::{draw_top, draw_input};
use crate::state::AppState;

pub fn  draw_game(ui: &mut Ui, state: &AppState) {
    panel(ui, Panel::top("top"), |ui| draw_top(ui, state));
    panel(ui, Panel::bottom("bottom"), |ui| draw_input(ui, state));
    //draw_room(ui, state);
    //draw_side(ui, state);
}

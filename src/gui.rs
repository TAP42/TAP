mod panels;
use panels::{draw_top, draw_input, draw_room, draw_side};
use eframe::egui;
use crate::state::AppState;


pub fn draw(ui: &mut egui::Ui, state: &AppState) {
    draw_top(ui, state);
    //draw_input(ui, state);
    //draw_room(ui, state);
    //draw_side(ui, state);
}

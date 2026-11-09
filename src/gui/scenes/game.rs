use eframe::egui::Ui;
use crate::gui::draw_top;
use crate::state::AppState;

pub fn  draw_game(ui: &mut Ui, state: &AppState) {
    draw_top(ui, state);
    //draw_input(ui, state);
    //draw_room(ui, state);
    //draw_side(ui, state);
}

pub mod panels;
pub mod scenes;
pub mod components;

use eframe::egui;
use crate::state::AppState;
use crate::gui::panels::{draw_top};
use crate::gui::scenes::{draw_menu, draw_game};

pub fn draw(ui: &mut egui::Ui, state: &mut AppState) {
    draw_menu(ui, state);
}

pub mod panels;
pub mod scenes;
pub mod components;

use eframe::egui;
use crate::state::{AppState, Scene};
use crate::gui::panels::{draw_top};
use crate::gui::scenes::{draw_menu, draw_game, byebye, settings};

pub fn draw(ui: &mut egui::Ui, state: &mut AppState) {
    match state.scene {
        Scene::Menu => draw_menu(ui, state),
        Scene::Game => draw_game(ui, state),
        Scene::Quit => byebye(ui, state),
        Scene::Settings => settings(ui, state),
    }
}

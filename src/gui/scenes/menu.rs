//menu with buttons: start, settings, quit
use super::game::draw_game;
use eframe::egui::{Ui, RichText};
use crate::gui::components::{button};
use crate::state::{AppState, Scene};
use crate::palette::{colored, sized, bold, *};

pub fn  draw_menu(ui: &mut Ui, state: &mut AppState) {
    let text_size: f32 = 20.0;

    ui.add_space((ui.available_height() - 80.0) / 2.0);
    ui.vertical_centered(|ui| {
        if button(ui, sized(bold("START"), text_size)).clicked() {
            state.scene = Scene::Game;
        }

        if button(ui, sized(bold("SETTINGS"), text_size)).clicked() {
            //draw_settings();
            todo!();
        }

        if button(ui, sized(bold("QUIT"), text_size)).clicked() {
            //exit();
            todo!();
        }
    });
}

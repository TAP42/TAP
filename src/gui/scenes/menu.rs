//menu with buttons: start, settings, quit
use eframe::egui::Ui;
use super::lobby::connect;
use crate::gui::components::{button, centered};
use crate::state::{AppState, Scene};
use crate::palette::{colored, sized, bold, *};

pub fn  draw_menu(ui: &mut Ui, state: &mut AppState) {
    let text_size: f32 = state.settings.text_size;

    ui.add_space((ui.available_height() - 80.0) / 2.0);
    centered(ui, |ui| {
        if button(ui, &state.settings, sized(bold("START"), text_size)).clicked() {
            state.show_connect = true;
            state.go_to(Scene::Lobby);
        }

        if button(ui, &state.settings, sized(bold("SETTINGS"), text_size)).clicked() {
            state.go_to(Scene::Settings);
        }

        if button(ui, &state.settings, sized(bold("QUIT"), text_size)).clicked() {
            state.go_to(Scene::Quit);
        }
    });
}

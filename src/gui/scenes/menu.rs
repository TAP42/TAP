//menu with buttons: start, settings, quit
use eframe::egui::Ui;
use crate::gui::components::{button};
use crate::state::{AppState, Scene};
use crate::palette::{colored, sized, bold, *};

pub fn  draw_menu(ui: &mut Ui, state: &mut AppState) {
    let text_size: f32 = state.settings.text_size;

    ui.add_space((ui.available_height() - 80.0) / 2.0);
    ui.vertical_centered(|ui| {
        if button(ui, state, sized(bold("START"), text_size)).clicked() {
            state.go_to(Scene::Game);
        }

        if button(ui, state, sized(bold("SETTINGS"), text_size)).clicked() {
            state.go_to(Scene::Settings);
        }

        if button(ui, state, sized(bold("QUIT"), text_size)).clicked() {
            state.go_to(Scene::Quit);
        }
    });
}

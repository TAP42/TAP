//settings with text changing, themes, etc.
use eframe::egui::Ui;
use crate::gui::components::back;
use crate::state::{AppState, Scene, Settings};
use crate::palette::{colored, sized, bold, *};

pub fn  settings(ui: &mut Ui, state: &mut AppState) {
    let text_size: f32 = state.settings.text_size;

    ui.vertical_centered(|ui| {
        ui.label(sized(bold("This will be implemented later."), text_size));
        ui.label(sized(bold("Here's a button to go back."), text_size));
        back(ui, state);
    });
}

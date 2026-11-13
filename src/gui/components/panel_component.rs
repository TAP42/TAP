use eframe::egui::{Ui, RichText, ProgressBar};
use crate::state::AppState;
use crate::palette::{colored, sized, bold, *};

pub fn centered<R>(ui: &mut Ui, add_contents: impl FnOnce(&mut Ui) -> R) {
    ui.vertical_centered(|ui| {
        ui.add_space(8.0);
        add_contents(ui);
    });
}

pub fn row<R>(ui: &mut Ui, add_contents: impl FnOnce(&mut Ui) -> R) {
    ui.horizontal(|ui| {
        ui.spacing_mut().item_spacing.x = 10.0;
        add_contents(ui)
    });
}

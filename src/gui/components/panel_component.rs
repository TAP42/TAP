use eframe::egui::{Ui, RichText, ProgressBar, Align, Layout, Panel, Frame};
use crate::state::AppState;
use crate::palette::{colored, sized, bold, *};

pub fn  panel<R>(ui: &mut Ui, panel: Panel, add_contents: impl FnOnce(&mut Ui) -> R) {
    let frame = Frame::new()
        .fill(ui.visuals().panel_fill)
        .stroke(ui.visuals().widgets.noninteractive.bg_stroke)
        .inner_margin(8.0);
    panel.frame(frame).show(ui, add_contents);
}

pub fn  chat<R>(ui: &mut Ui, panel: Panel, add_contents: impl FnOnce(&mut Ui) -> R) {
    let frame = Frame::new()
        .fill(ui.visuals().panel_fill)
        .stroke(ui.visuals().widgets.noninteractive.bg_stroke);
    panel.frame(frame).show(ui, add_contents);
}

pub fn  centered<R>(ui: &mut Ui, add_contents: impl FnOnce(&mut Ui) -> R) {
    ui.vertical_centered(|ui| {
        ui.add_space(8.0);
        add_contents(ui);
    });
}

pub fn  row<R>(ui: &mut Ui, add_contents: impl FnOnce(&mut Ui) -> R) {
    ui.horizontal(|ui| {
        ui.spacing_mut().item_spacing.x = 10.0;
        add_contents(ui)
    });
}

pub fn  align<R>(ui: &mut Ui, alignment: &Align, add_contents: impl FnOnce(&mut Ui) -> R) {
    ui.with_layout(Layout::right_to_left(*alignment), |ui| {
        add_contents(ui);
    });
}

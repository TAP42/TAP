use eframe::egui::{Ui, Window, Align2};
use crate::state::{AppState, Scene};

pub fn  window<R>(
    ui: &mut Ui,
    title: &str,
    open: &mut bool,
    add_contents:
    impl FnOnce(&mut Ui) -> R
    ) {
    Window::new(title)
        .open(open)
        .collapsible(false)
        .resizable(false)
        .pivot(Align2::CENTER_CENTER)
        .default_pos(ui.content_rect().center())
        .show(ui, add_contents);
}

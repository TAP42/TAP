use eframe::egui::{Ui, Response, TextEdit, WidgetText};
use crate::palette::{colored, sized, bold, *};
use crate::state::{Settings, AppState};

pub fn  field(ui: &mut Ui, settings: &Settings, text: impl Into<WidgetText>, var: &mut String) -> Response {
    let w: f32 = 150.0 * settings.text_size;
    let h: f32 = 10.0 * settings.text_size;

    ui.label(text);
    ui.add_sized([w, h], TextEdit::singleline(var))
}

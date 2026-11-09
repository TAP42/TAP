use eframe::egui::{Ui, Response, Button, WidgetText};
use crate::palette::{colored, sized, bold, *};

pub fn  button(ui: &mut Ui, text: impl Into<WidgetText>) -> Response {
    //add_sized([w, h]): forces a fixed size
    ui.add_sized([150.0, 60.0], Button::new(text))
}

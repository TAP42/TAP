use eframe::egui::{Ui, Response, Button, WidgetText};
use crate::palette::{colored, sized, bold, *};
use crate::state::AppState;

pub fn  button(ui: &mut Ui, state: &mut AppState, text: impl Into<WidgetText>) -> Response {
    //add_sized([w, h]): forces a button size
    let w: f32 = 120.0 * state.settings.button_w;
    let h: f32 = 60.0 * state.settings.button_h;

    ui.add_sized([w, h], Button::new(text))
}

pub fn  back(ui: &mut Ui, state: &mut AppState) {
    if button(ui, state, bold("RETURN")).clicked() {
        state.go_back();
    }
}

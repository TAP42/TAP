use eframe::egui::{Ui, Response, WidgetText, Slider, SliderClamping};
use std::ops::RangeInclusive;
use crate::palette::{colored, sized, bold, *};

pub fn  slider(ui: &mut Ui, label: impl Into<WidgetText>, value: &mut f32, range: impl Into<RangeInclusive<f32>>) -> Response {
    ui.add(Slider::new(value, range)
    .text(label)
    .clamping(SliderClamping::Always)
    .step_by(0.05))
}

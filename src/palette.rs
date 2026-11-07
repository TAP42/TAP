use eframe::egui::{Color32, RichText};

pub const   RED: Color32 = Color32::RED;
pub const   GREEN: Color32 = Color32::GREEN;
pub const   YELLOW: Color32 = Color32::YELLOW;

// use colored(text, any of the const above), todo add more colours

pub fn  colored(text: impl Into<RichText>, color: Color32) -> RichText {
    text.into().color(color)
}

// use sized(text, float value)

pub fn  sized(text: impl Into<RichText>, size: f32) -> RichText {
    text.into().size(size)
}

// use bold(text)

pub fn  bold(text: impl Into<RichText>) -> RichText {
    text.into().strong()
}

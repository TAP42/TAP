use eframe::egui::{self, Color32, RichText, Context, Visuals};
use crate::state::ThemeKind;

pub const   RED: Color32 = Color32::RED;
pub const   GREEN: Color32 = Color32::GREEN;
pub const   YELLOW: Color32 = Color32::YELLOW;

pub const   GREY: Color32 = Color32::GRAY;
pub const   D_GREY: Color32 = Color32::DARK_GRAY;
pub const   L_GREY: Color32 = Color32::LIGHT_GRAY;

pub const   WHITE: Color32 = Color32::WHITE;
pub const   BLACK: Color32 = Color32::BLACK;

// below is the theme stuff. todo add custom theming via 'themes/', later.

impl    From<ThemeKind> for Theme {
    fn  from(kind: ThemeKind) -> Self {
        match kind {
            ThemeKind::Dark => Theme::dark(),
            ThemeKind::Light => todo!(),
        }
    }
}

pub struct  Theme {
    pub background: Color32,
    pub panel: Color32,
    pub widget: Color32,
    pub accent: Color32,
    pub text: Color32,
}

impl    Theme {
    pub fn  dark() -> Self {
        Self {
            background: Color32::from_rgb(30, 30, 30),
            panel: Color32::from_rgb(20, 20, 20),
            widget: Color32::from_rgb(40, 40, 40),
            accent: Color32::from_rgb(150, 120, 60),
            text: Color32::from_rgb(240, 240, 240),
        }
    }

    pub fn  apply(&self, ctx: &Context) {
        let mut v = Visuals::dark();
        v.panel_fill = self.panel;
        v.window_fill = self.background;
        v.extreme_bg_color = self.background;
        v.override_text_color = Some(self.text);
        v.selection.bg_fill = self.accent;
        v.widgets.inactive.bg_fill = self.widget;
        v.widgets.inactive.weak_bg_fill = self.widget;
        ctx.set_visuals(v);
    }
}

// use colored(text, any of the const above), todo add more colours

pub fn  colored(text: impl Into<RichText>, color: Color32) -> RichText {
    text.into().color(color)
}

// use sized(text, float value)

pub fn  sized(text: impl Into<RichText>, size: f32) -> RichText {
    text.into().size(20.0 * size)
}

// use bold(text)

pub fn  bold(text: impl Into<RichText>) -> RichText {
    text.into().strong()
}

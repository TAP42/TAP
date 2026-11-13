use serde::{Deserialize, Serialize};
use serde_json::{from_str, to_string_pretty};
use std::fs::{read_to_string, write};
use eframe::egui::Context;

#[derive(Default, PartialEq, Clone, Copy)]
pub enum    Scene {
    #[default]
    Menu,
    Game,
    Quit,
    Settings,
}

const   MIN_SCALE: f32 = 0.5;
const   MAX_SCALE: f32 = 2.0;
const   SETTINGS_PATH: &str = "settings.json";

#[derive(Serialize, Deserialize, Clone, Copy, PartialEq)]
pub enum    ThemeKind {
    Dark,
    Light,
}

#[derive(Serialize, Deserialize, Clone)]
#[serde(default)]
pub struct  Settings {
    pub theme: ThemeKind,
    pub text_size: f32,
    pub button_w: f32,
    pub button_h: f32,
}

impl    Default for Settings {
    fn  default() -> Self {
        //defaults
        Self {
            theme: ThemeKind::Dark,
            text_size: 1.0,
            button_w: 1.0,
            button_h: 1.0,
        }
    }
}

impl    Settings {
    pub fn  sanitize(&mut self) {
        //checks whether the values are within the range MIN_SCALE to MAX_SCALE
        let d = Settings::default();
        for (v, fallback) in [
            (&mut self.text_size, d.text_size),
            (&mut self.button_w, d.button_w),
            (&mut self.button_h, d.button_h),
        ] {
            *v = if v.is_finite() { v.clamp(MIN_SCALE, MAX_SCALE) } else { fallback };
        }
    }

    pub fn  load() -> Self {
        //loads settings.json, if settings.json doesnt exist or is faulty, default
        let mut s: Settings = read_to_string(SETTINGS_PATH)
            .ok()
            .and_then(|txt| from_str(&txt).ok())
            .unwrap_or_default();
        s.sanitize();
        s
    }

    pub fn  save(&self) {
        //saves current settings to settings.json
        if let Ok(txt) = to_string_pretty(self) {
            let _ = write(SETTINGS_PATH, txt);
        }
    }
}

pub struct  PlayerStats {
    pub max_hp: i32,
    pub hp: i32,
    pub atk: i32,
    pub def: i32,
    pub spd: i32,
    pub name: String,
}

impl Default for PlayerStats {
    fn  default() -> Self {
        Self {
            max_hp: 100,
            hp: 100,
            atk: 1,
            def: 1,
            spd: 1,
            name: "Name".to_string(),
        }
    }
}

#[derive(Default)]
pub struct  AppState {
    pub scene: Scene,
    pub previous_scene: Option<Scene>,
    pub stats: PlayerStats,
    pub connected: bool,
    pub host: String,
    pub port: String,
    pub selected_channel: i32,
    pub input_text: String,
    pub chat: Vec<Vec<String>>,
    pub logs: Vec<String>,
    pub settings: Settings,
}

impl    AppState {
    pub fn  go_to(&mut self, scene: Scene) {
        if scene != self.scene {
            self.previous_scene = Some(self.scene);
            self.scene = scene;
        }
    }

    pub fn  go_back(&mut self) {
        if let Some(prev) = self.previous_scene.take() {
            self.scene = prev;
        }
    }
}

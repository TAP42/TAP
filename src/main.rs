mod gui;
mod palette;
mod state;
mod network;

use gui::draw;
use eframe::{egui, Result};
use crate::state::{Settings, ThemeKind};
use crate::palette::Theme;
use crate::gui::scenes::settings;

#[derive(Default)]
struct  TApp {
    state: state::AppState,
}

impl eframe::App for TApp {
    fn ui(&mut self, ui: &mut egui::Ui, _frame: &mut eframe::Frame) {
        self.state.poll_network();
        egui::CentralPanel::default().show(ui, |ui| {
            draw(ui, &mut self.state)
        });
    }
}

fn main() -> eframe::Result {
    let options = eframe::NativeOptions::default();
    eframe::run_native(
        "testing",
        options,
        Box::new(|cc| {
            let settings = Settings::load();
            Theme::from(settings.theme).apply(&cc.egui_ctx);
            let mut app = TApp::default();
            app.state.settings = settings;
            Ok(Box::new(app))
        }),
    )
}

mod gui;
mod palette;
mod state;
use gui::draw;
use eframe::egui;

#[derive(Default)]
struct  TApp {
    app: state::AppState,
}

impl eframe::App for TApp {
    fn ui(&mut self, ui: &mut egui::Ui, _frame: &mut eframe::Frame) {
        egui::CentralPanel::default().show(ui, |ui| {
            ui.heading("testing");
            draw(ui, &self.app)
        });
    }
}

fn  main() {
    let options = eframe::NativeOptions::default();
    let _ = eframe::run_native(
        "testing",
        options,
        Box::new(|_cc| Ok(Box::new(TApp::default()))),
    );
}

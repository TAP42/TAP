mod gui;
use gui::counter_gui;
use eframe::egui;

#[derive(Default)]
struct  TApp {
    counter: i32,
}

impl eframe::App for TApp {
    fn ui(&mut self, ui: &mut egui::Ui, _frame: &mut eframe::Frame) {
        egui::CentralPanel::default().show(ui, |ui| {
            ui.heading("counter");
            counter_gui(ui, &mut self.counter);
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

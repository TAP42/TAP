use eframe::egui::{Ui, RichText, ProgressBar};
use crate::state::AppState;
use crate::palette::{colored, sized, bold, *};

pub fn draw_template(ui: &mut Ui, state: &AppState) {
    // adding a paragraph
    ui.horizontal(|ui| {
        // adding text
        ui.label("text");
        // text sizing 10px
        ui.label(sized("text", 10.0));
        // bold text
        ui.label(bold("text"));
        // text colored in red
        ui.label(colored("text", RED));
        // separating
        ui.separator();
    });

    // to add a progress bar
    let progress: f32 = 20.0;
    ui.add(ProgressBar::new(progress).show_percentage();
}

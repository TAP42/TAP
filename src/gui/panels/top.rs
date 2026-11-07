// top with player stats, player name, server connection status
use eframe::egui::{Ui, RichText, ProgressBar};
use crate::state::AppState;
use crate::palette::{colored, sized, bold, *};

fn  connection_code(c: &bool) -> RichText {
    match c {
        true    => colored("Connected", GREEN),
        false   => colored("Disconnected", RED),
    }
}

pub fn draw_top(ui: &mut Ui, state: &AppState) {
    let progress = state.stats.hp;

    ui.horizontal(|ui| {
        ui.label(sized(bold("Status:"), 20.0));
        ui.label(sized(bold(connection_code(&state.connected)), 20.0));
        ui.separator();
        ui.label(sized(format!("{} - {}/{}HP",
            state.stats.name,
            state.stats.hp,
            state.stats.max_hp), 20.0));
    });
    ui.add(ProgressBar::new(progress as f32).show_percentage().desired_width(400.0).desired_height(15.0));
    ui.separator();
}

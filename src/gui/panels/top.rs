// top with player stats, player name, server connection status
use eframe::egui::{Ui, RichText, ProgressBar};
use crate::gui::components::row;
use crate::state::AppState;
use crate::palette::{colored, sized, bold, *};

fn  connection_code(c: &bool) -> RichText {
    match c {
        true    => colored("Connected", GREEN),
        false   => colored("Disconnected", RED),
    }
}

pub fn draw_top(ui: &mut Ui, state: &AppState) {
    let text_size: f32 = state.settings.text_size;
    let progress = state.stats.hp;

    row(ui, |ui| {
        ui.label(sized(bold("Status:"), text_size));
        ui.label(sized(bold(connection_code(&state.connected)), text_size));
        ui.separator();
        ui.label(sized(format!("{} - {}/{}HP",
            state.stats.name,
            state.stats.hp,
            state.stats.max_hp), text_size));
        ui.separator();
        ui.label(sized(format!("Balance: {} tokens", state.stats.bal), text_size));
    });
    ui.add(ProgressBar::new(progress as f32).show_percentage().desired_width(400.0).desired_height(15.0));
    ui.separator();
}

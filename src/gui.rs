use eframe::egui;


pub fn  counter_gui(ui: &mut egui::Ui, counter: &mut i32) {
    let temp: &str = "https://github.com/TAP42/TAP";

    // .horizontal: une ligne/un groupe de components
    // .button, .label: components bouton et text field

    ui.horizontal(|ui| {
        if ui.button("+").clicked() {
            *counter += 1;
        }
        ui.label(counter.to_string());
        ui.hyperlink(temp);
    });
}

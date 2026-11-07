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
    pub stats: PlayerStats,
    pub connected: bool,
    pub host: String,
    pub port: String,
    pub selected_channel: i32,
    pub input_text: String,
    pub chat: Vec<Vec<String>>,
    pub logs: Vec<String>,
}


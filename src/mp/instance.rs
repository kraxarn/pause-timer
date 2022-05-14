use std::time::Duration;
use dbus::blocking::Connection;
use crate::mp::MediaPlayerInstance;

impl MediaPlayerInstance {
	fn from_name(name: &str, connection: &Connection) -> Option<Self> {
		let proxy = connection.with_proxy(format!("org.mpris.MediaPlayer2.{}", name),
			"/", Duration::from_secs(3));
	}
}
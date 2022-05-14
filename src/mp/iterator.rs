use std::ops::Index;
use std::time::Duration;
use dbus::blocking::Connection;
use super::MediaPlayerInstance;
use super::MediaPlayerIterator;

impl<'a> MediaPlayerIterator<'a> {
	fn new(connection: &'a Connection) -> Result<Self, std::io::Error> {
		let proxy = connection.with_proxy("org.freedesktop.DBus", "/", Duration::from_secs(3));
		let (names, ) = proxy.method_call("org.freedesktop.DBus", "ListNames", ())?;

		Ok(MediaPlayerIterator {
			connection,
			names,
			current: 0_usize,
		})
	}
}

impl<'a> Iterator for MediaPlayerIterator<'a> {
	type Item = MediaPlayerInstance;

	fn next(&mut self) -> Option<Self::Item> {
		if self.current + 1 >= self.names.len() {
			return None;
		}

		loop {
			if self.names[self.current].starts_with("org.mpris.MediaPlayer2.") {

			}

			self.current += 1;
		}

		None
	}
}

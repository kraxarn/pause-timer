use dbus::blocking::Connection;

mod iterator;
mod instance;

struct MediaPlayerIterator<'a> {
	connection: &'a Connection,
	names: Vec<String>,
	current: usize,
}

pub struct MediaPlayerInstance {
	name: String,
	identity: String,
}
// ==UserScript==
// @name         PauseTimer
// @homepage     https://github.com/kraxarn/pause-timer
// @version      1.1.0
// @encoding     utf-8
// @author       kraxarn
// @match        *://*.youtube.com/*
// @exclude      *://music.youtube.com/*
// @exclude      *://*.music.youtube.com/*
// @compatible   firefox
// @downloadURL  https://github.com/kraxarn/pause-timer/blob/master/pause-timer.user.js
// @updateURL    https://github.com/kraxarn/pause-timer/blob/master/pause-timer.user.js
// @run-at       document-end
// ==/UserScript==

enum PlayerState {
	Playing = 1,
	Paused = 2,
	Loading = 3,
}

interface VideoPlayer {
	pauseVideo(): void

	playVideo(): void

	getPlayerState(): PlayerState
}

class PauseTimer {
	private player: VideoPlayer
	private running = false

	constructor() {
		document.onclick = event => this.onClick(event)
	}

	set status(value: string) {
		const infoStrings = document.querySelector("#description #info")
		if (!infoStrings) {
			return
		}

		let status = document.getElementById("pause-timer-status")
		if (!status) {
			const spacer = document.createElement("span")
			spacer.textContent = "  "
			infoStrings.appendChild(spacer)

			status = document.createElement("span")
			status.id = "pause-timer-status"
			infoStrings.appendChild(status)
		}

		status.textContent = value
	}

	private onClick(event: MouseEvent) {
		const target = <HTMLElement>event.target;
		if (target.id !== "expand"
			|| target.parentElement.id !== "description-inline-expander"
			|| document.getElementById("pause-timer-container")) {
			return
		}

		const container = document.createElement("div")
		container.id = "pause-timer-container"
		container.style.display = "flex"
		container.style.flexDirection = "row"
		container.style.gap = "10px"
		container.style.paddingTop = "10px"

		const title = document.createElement("span")
		title.textContent = "Pause timer"
		container.appendChild(title)

		const minutes = document.createElement("input")
		minutes.type = "number"
		minutes.value = "5"
		container.appendChild(minutes)

		const apply = document.createElement("input")
		apply.type = "button"
		apply.value = "Apply"
		apply.onclick = () => this.onApply(parseInt(minutes.value))
		container.appendChild(apply)

		const cancel = document.createElement("input")
		cancel.type = "button"
		cancel.value = "Cancel"
		cancel.onclick = () => this.running = false
		container.appendChild(cancel)

		document
			.querySelector("#description-inline-expander .ytd-text-inline-expander > .yt-formatted-string:last-child")
			.appendChild(container)
	}

	private async waitForState(state: PlayerState): Promise<void> {
		return new Promise(resolve => {
			const id = setInterval(() => {
				if (this.player.getPlayerState() === state) {
					clearInterval(id)
					resolve()
				}
			})
		})
	}

	private async sleep(timeout: number): Promise<void> {
		return new Promise(resolve => setTimeout(() => resolve(), timeout))
	}

	private async onApply(minutes: number) {
		if (!this.player) {
			alert("Error: no player found")
			return
		}

		this.running = true
		while (this.running) {
			this.status = "Waiting..."
			await this.waitForState(PlayerState.Playing)

			for (let i = 0; i < minutes; i++) {
				this.status = `Waiting (${i + 1}/${minutes})...`
				await this.sleep(60_000)
			}
			this.player.pauseVideo()
		}
	}

	public onPlayerUpdated(event: CustomEvent) {
		this.player = event.detail
	}
}

let pauseTimer: PauseTimer;
window.addEventListener("yt-navigate-finish", () => pauseTimer = new PauseTimer(), true)
window.addEventListener("yt-player-updated", (e: CustomEvent) => pauseTimer.onPlayerUpdated(e))
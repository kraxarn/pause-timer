// ==UserScript==
// @name         PauseTimer
// @homepage     https://github.com/kraxarn/pause-timer
// @version      1.0.5
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
var PlayerState;
(function (PlayerState) {
    PlayerState[PlayerState["Playing"] = 1] = "Playing";
    PlayerState[PlayerState["Paused"] = 2] = "Paused";
    PlayerState[PlayerState["Loading"] = 3] = "Loading";
})(PlayerState || (PlayerState = {}));
class PauseTimer {
    constructor() {
        document.onclick = event => this.onClick(event);
    }
    set status(value) {
        const infoStrings = document.querySelector("#description #info");
        if (!infoStrings) {
            return;
        }
        let status = document.getElementById("pause-timer-status");
        if (!status) {
            const spacer = document.createElement("span");
            spacer.textContent = "  ";
            infoStrings.appendChild(spacer);
            status = document.createElement("span");
            status.id = "pause-timer-status";
            infoStrings.appendChild(status);
        }
        status.textContent = value;
    }
    onClick(event) {
        const target = event.target;
        if (target.id !== "expand"
            || target.parentElement.id !== "description-inline-expander"
            || document.getElementById("pause-timer-container")) {
            return;
        }
        const container = document.createElement("div");
        container.id = "pause-timer-container";
        container.style.display = "flex";
        container.style.flexDirection = "row";
        container.style.gap = "10px";
        container.style.paddingTop = "10px";
        const title = document.createElement("span");
        title.textContent = "Pause timer";
        container.appendChild(title);
        const minutes = document.createElement("input");
        minutes.type = "number";
        minutes.value = "5";
        container.appendChild(minutes);
        const apply = document.createElement("input");
        apply.type = "button";
        apply.value = "Apply";
        apply.onclick = () => this.onApply(parseInt(minutes.value));
        container.appendChild(apply);
        document
            .querySelector("#description-inline-expander .ytd-text-inline-expander > .yt-formatted-string:last-child")
            .appendChild(container);
    }
    log(...message) {
        console.log("[PauseTimer]", ...message);
    }
    async waitForState(state) {
        return new Promise(resolve => {
            const id = setInterval(() => {
                if (this.player.getPlayerState() === state) {
                    clearInterval(id);
                    resolve();
                }
            });
        });
    }
    async sleep(timeout) {
        return new Promise(resolve => setTimeout(() => resolve(), timeout));
    }
    async onApply(minutes) {
        if (!this.player) {
            alert("Error: no player found");
            return;
        }
        while (true) {
            this.status = "Waiting...";
            await this.waitForState(PlayerState.Playing);
            for (let i = 0; i < minutes; i++) {
                this.status = `Waiting (${i + 1}/${minutes})...`;
                await this.sleep(60000);
            }
            this.player.pauseVideo();
        }
    }
    onPlayerUpdated(event) {
        this.player = event.detail;
    }
}
let pauseTimer;
window.addEventListener("yt-navigate-finish", () => pauseTimer = new PauseTimer(), true);
window.addEventListener("yt-player-updated", (e) => pauseTimer.onPlayerUpdated(e));

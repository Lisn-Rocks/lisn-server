<template>
    <section class="player-minimized" v-bind:class="{ shown: isShown }">
        <div class="progress-bar">
            <div class="progress" v-bind:style="{ width: progress * 100 + '%' }"/>
        </div>

        <div class="song">
            <div class="song-info">
                <h4 class="song-title">{{ currentSongInfo.name }}</h4>
                <h6 class="song-artist">{{ currentSongInfo.artist }}</h6>
            </div>

            <div class="song-control" v-on:click="$emit('toggle')">
                <i class="material-icons">{{ action }}_circle_filled</i>
            </div>
        </div>
    </section>
</template>

<script>
export default {
    name: 'PlayerMin',

    props: {
        isShown: Boolean,
        progress: Number,
        paused: Boolean,
        currentSongInfo: Object,
    },

    data() {
        return {
            action: this.paused? "play" : "pause"
        }
    }
}
</script>

<style scoped>
.player-minimized {
    position: fixed;
    width: 100%;
    height: var(--player-minimized-height);
    bottom: 0;
    left: 0;
    color: white;
    /* cursor: pointer; */
    background-color: var(--main-dark-color);
    display: none;
}

.player-minimized .progress-bar {
    width: 100%;
    height: 3px;
    background-color: var(--secondary-white);
}
.player-minimized .progress-bar .progress {
    height: 100%;
    background-color: var(--accent-color);
}

.player-minimized .song {
    display: flex;
    height: var(--player-minimized-height);
    color: white;
    padding: 0 var(--standard-horizontal-padding);
}

.player-minimized .song .song-info {
    flex-grow: 2;
    display: flex;
    flex-direction: column;
    justify-content: center;
    overflow: hidden;
}
.player-minimized .song .song-info .song-title {
    font-size: var(--song-title-font-size);
    font-weight: 500;
}
.player-minimized .song .song-info .song-artist {
    font-size: var(--song-artist-font-size);
    font-weight: 400;
    color: var(--secondary-white);
    line-height: 1.1em;
}

.player-minimized .song .song-control {
    cursor: pointer;
    padding: 0 10px;
    margin-right: -15px;
}
.player-minimized .song .song-control i {
    line-height: var(--player-minimized-height);
    font-size: 32px;
}

.shown {
    display: block !important;
}
</style>

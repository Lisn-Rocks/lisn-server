<template>
    <section class="player-minimized" v-bind:class="{ shown: isShown }">
        <ProgressBarMin v-bind:currentSong="currentSong" :key="b"/>

        <div class="song">
            <div class="album-cover" v-bind:style="coverStyle"/>

            <div class="song-info">
                <h4 class="song-title">{{ currentSongInfo.name }}</h4>
                <h6 class="song-artist">{{ currentSongInfo.artist }}</h6>
            </div>

            <div class="song-control" v-on:click="$emit('toggle')">
                <i class="material-icons">{{ action }}</i>
            </div>
        </div>
    </section>
</template>

<script>
import ProgressBarMin from './ProgressBarMin.vue';

export default {
    name: 'PlayerMin',

    components: {
        ProgressBarMin
    },

    data() {
        return {
            b: false
        }
    },

    props: {
        isShown: Boolean,
        currentSong: HTMLAudioElement,
        currentSongInfo: Object,
        currentAlbumCoverURL: String,
    },

    computed: {
        coverStyle: function() {
            return `background-image: url('${this.currentAlbumCoverURL}');`;
        },

        action: function() {
            return this.currentSong.paused? 'play_arrow' : 'pause';
        }
    },

    created() {
        this.currentSong.addEventListener(
            'timeupdate', () => this.progressBarUpdate()
        );
    },

    methods: {
        progressBarUpdate: function() {
            this.b = !this.b;
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

.player-minimized .song {
    display: flex;
    height: var(--player-minimized-height);
    color: white;
}

.player-minimized .song .album-cover {
    width: var(--player-minimized-height);
    height: var(--player-minimized-height);
    margin-right: 10px;
    background-repeat: no-repeat;
    background-size: cover;
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
    padding: 0 10px 0 var(--standard-horizontal-padding);
}
.player-minimized .song .song-control i {
    line-height: var(--player-minimized-height);
    font-size: 32px;
}

.shown {
    display: block !important;
}
</style>

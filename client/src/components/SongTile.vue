<template>
    <a class="song-tile"
        v-bind:class="{ 'song-active': isPlaying }"
        v-on:click="$emit('play')"
    >
        <a class="song-info">
            <h4 class="song-title">{{ song.title }}</h4>
            <h6 class="song-artist">{{ song.artist }}</h6>
        </a>

        <p class="song-duration">{{ song.minutes }}:{{ song.seconds }}</p>
    </a>
</template>

<script>
export default {
    name: "SongTile",
    props: {
        song: Object, 
        isPlaying: Boolean
    }
}
</script>

<style scoped>
.song-tile {
    display: flex;
    height: var(--song-tile-height);
    color: white;
    padding: 0 var(--standard-horizontal-padding);
    cursor: pointer;
    transition: all 0.2s;
}

.song-tile .song-info {
    flex-grow: 2;
    display: flex;
    flex-direction: column;
    justify-content: center;
}
.song-tile .song-info .song-title {
    font-size: var(--song-title-font-size);
    font-weight: 500;
    line-height: 1.4em;
}
.song-tile .song-info .song-artist {
    font-size: var(--song-artist-font-size);
    font-weight: 400;
    color: var(--secondary-white);
}

.song-tile .song-duration {
    line-height: var(--song-tile-height);
    font-weight: 300;
    font-size: var(--song-duration-font-size);
}

.song-active {
    color: var(--accent-color);
    transition: all 0.2s;
}

.song-active .song-info .song-title {
    font-weight: 700;
}
.song-active .song-info .song-artist {
    color: var(--accent-light-color);
    font-weight: 500;
}

.song-active .song-duration {
    font-weight: 400;
}
</style>
<template>
    <div id="queue" v-bind:class="{ shown: isShown }">
        <div v-bind:key="song.id" v-for="song in queue">
            <SongTile
                v-bind:song="song"
                v-bind:isPlaying="currentSongID === song.id"
                v-on:play="$emit('playSong', song.id)"
            />
        </div>
    </div>
</template>

<script>
import SongTile from './SongTile.vue'

export default {
    name: 'Queue',
    components: {
        SongTile
    },
    props: {
        isShown: Boolean,
        queue: Array,
        currentSongID: Number
    }
}
</script>

<style scoped>
#queue {
    display: none;
}

.shown {
    display: block !important;
}
</style>

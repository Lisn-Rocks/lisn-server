<template>
    <div id="queue" v-bind:class="{ shown: isShown }">
        <div v-bind:key="song.qid" v-for="song in queue">
            <SongTile
                v-bind:song="song"
                v-bind:isPlaying="currentSongQID === song.qid"
                v-on:play="$emit('goto', song.qid)"
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
        currentSongQID: Number
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

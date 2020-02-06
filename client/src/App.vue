<template>
    <div id="app">
        <Topbar v-bind:tabs="tabs"/>
        <main>
            <section v-bind:key="song.id" v-for="song in queue">
                <SongTile v-bind:song="song" v-bind:songs="queue"/>
            </section>
        </main>
    </div>
</template>

<script>
import Topbar from './components/Topbar/Topbar.vue'
import SongTile from './components/SongTile/SongTile.vue'

export default {
    name: 'app',

    components: {
        Topbar,
        SongTile
    },

    data() {
        return {
            tabs: [
                {
                    name: "Queue",
                    isActive: true
                },
                {
                    name: "Search",
                    isActive: false
                }
            ],

            // The song queue is fetched from server by the getQueue method.
            queue: []
        }
    },

    created() {
        this.fetchQueue();
    },

    methods: {
        fetchQueue() {
            for (let i = 1; i < 5; i++) {
                fetch('http://192.168.0.24:8000/songinfo/' + i)
                    .then( response => response.json() )
                    .then( song => {
                        song.isActive = false;
                        song.minutes = Math.floor(song.duration / 60).toString();
                        song.seconds = (song.duration % 60).toString();
                        song.seconds = (song.seconds.length < 2) ? '0' + song.seconds : song.seconds;
                        this.queue.push(song);
                    } );
            }
        }
    }
}
</script>

<style>
* {
    margin: 0;
    padding: 0;
    outline: none;
}

:root {
    /* Scaling Factor (uncomment the proper one) */
    /* Windows */
    /* --scaling-factor: 0.64; */

    /* Linux | UNIX | Darwin */
    --scaling-factor: 0.64;

    /* Color Palette */
    --main-dark-color: #000A12;
    --main-color: #263238;
    --main-light-color: #4F5B62;

    --accent-color: #F50057;
    --accent-light-color: #FF5983;

    --secondary-white: #CCCCCC;

    /* General */
    --standard-horizontal-padding: 20px;

    /* Header */
    --header-height: 60px;
    --header-font-size: 20px;

    /* Main Section */
    --main-section-height: calc( 100vh - var(--header-height) );

    /* Song Tile */
    --song-tile-height: 75px;
    --song-title-font-size: 18px;
    --song-artist-font-size: 14px;
    --song-duration-font-size: 16px;
}

#app {
    font-family: 'Roboto', Helvetica, Arial, sans-serif;
    position: relative;
}

#app main {
    background-color: var(--main-color);
    padding-top: var(--header-height);
    height: var(--main-section-height);
}
</style>

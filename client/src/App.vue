<template>
    <div id="app">

        <header>
            <div v-bind:key="tab.id" v-for="tab in tabs">
                <TopbarTab 
                    v-bind:tab="tab"
                    v-bind:isActive="activeTabID === tab.id"
                    v-on:switch="activeTabID = tab.id"
                />
            </div>
        </header>

        <main>
            <Queue
                v-bind:queue="queue"
                v-bind:isShown="activeTabID === tabs.Queue.id"
                v-bind:currentSongID="currentSongID"
                v-on:playSong="fetchAndPlay($event)"
            />
            <Search v-bind:isShown="activeTabID === tabs.Search.id"/>
        </main>

        <PlayerMin 
            v-bind:isShown="currentSongID > 0"
            v-bind:currentSong="currentSong"
        />

    </div>
</template>


<script>
import TopbarTab from './components/TopbarTab.vue'
import Queue from './components/Queue.vue'
import Search from './components/Search.vue'
import PlayerMin from './components/PlayerMin.vue'

export default {
    name: 'app',

    components: {
        TopbarTab,
        Queue,
        Search,
        PlayerMin
    },

    data() {
        return {
            PROTO: 'http://',
            ROUTE: '10.14.199.118:8000', // must be changed appropriately

            tabs: {
                Queue: {
                    id: 0,
                    name: "Queue",
                },
                Search: {
                    id: 1,
                    name: "Search",
                }
            },

            activeTabID: 0,
            currentSongID: 0,
            currentSong: new Audio(),

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
                fetch(this.PROTO + this.ROUTE + '/songinfo/' + i)
                .then( response => response.json() )
                .then( song => {
                    song.isActive = false;
                    song.minutes = Math.floor(song.duration / 60).toString();
                    song.seconds = (song.duration % 60).toString();
                    song.seconds = (song.seconds.length < 2) 
                        ? '0' + song.seconds : song.seconds;
                    this.queue.push(song);
                } );
            }
        },

        async fetchAndPlay(songID) {
            this.currentSong.pause();

            this.currentSong = new Audio(this.PROTO + this.ROUTE + '/song/' + songID);
            this.currentSong.type = 'audio/mp3';

            this.currentSongID = songID;

            try {
                await this.currentSong.play();
            } catch (err) {
                alert('Failed to play: ' + err);
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

#app header {
    font-size: var(--header-font-size);
    position: fixed;
    width: 100%;
    top: 0;
    left: 0;
    height: var(--header-height);
    background-color: var(--main-light-color);
    display: flex;
    color: white;
    -webkit-box-shadow: 0px 5px 5px 0px rgba(0,0,0,0.5);
    -moz-box-shadow: 0px 5px 5px 0px rgba(0,0,0,0.5);
    box-shadow: 0px 5px 5px 0px rgba(0,0,0,0.5);
}

#app main {
    background-color: var(--main-color);
    padding-top: var(--header-height);
    height: var(--main-section-height);
}
</style>

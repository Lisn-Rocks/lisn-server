<template>
    <div id="app">
        <Topbar v-bind:tabs="tabs"/>
        <main>

        </main>
    </div>
</template>

<script>
import Topbar from './components/Topbar/Topbar.vue'

export default {
    name: 'app',
    components: {
        Topbar
    },
    data: () => {
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
            ]
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

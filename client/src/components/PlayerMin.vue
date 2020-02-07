<template>
    <button id="switch" 
        v-bind:class="{ shown: isShown }"
        v-on:click="toggle"
    >
        <i class="material-icons">{{ action }}_circle_filled</i>
    </button>
</template>

<script>
export default {
    name: 'PlayerMin',

    props: {
        isShown: Boolean,
        currentSong: HTMLAudioElement,
    },

    data() {
        return {
            action: "pause"
        }
    },

    methods: {
        toggle() {
            if (this.currentSong.ended) return;

            if (this.action === "play") {
                this.currentSong.play();
                this.action = "pause";
            } else {
                this.currentSong.pause();
                this.action = "play";
            }
        }
    }
}
</script>

<style scoped>
#switch {
    display: none;

    position: fixed;
    bottom: 20px;
    right: 20px;
    height: 60px;
    width: 60px;

    background-color: white;
    border: none;
    border-radius: 50%;
    cursor: pointer;
}

#switch i {
    color: var(--accent-color);
    font-size: 60px;
}

.shown {
    display: block !important;
}
</style>

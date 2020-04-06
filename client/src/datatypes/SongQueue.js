export function SongQueue() {

    // q contains songs in the order they were added.
    this.q = [];

    // repeat signifies the repeat mode of the SongQueue.
    // There are 3 such modes:
    //
    //     -1 = 'no repeat'
    //   * Finished or skipped songs get removed from SongQueue.
    //   * 'prev' automatically turns 'no repeat' to 'repeat queue' and plays
    //     the last song in the SongQueue.
    //
    //     0 = 'repeat queue'
    //   * Finished or skipped songs are pushed to the end of SongQueue.
    //
    //     1 = 'repeat one song'
    //   * 'prev' or 'next' force 'repeat one song' mode to turn to
    //     'repeat queue' and then do whatever they do.
    //
    this.repeat = 0;

    this.shuffle = false;

    this.get = function(qid) {
        return this.q[qid];
    }

    this.push = function(song) {
        song.qid = this.q.length;
        this.q.push(song);
    };

    this.prev = function() {
        alert("prev");
        switch (this.repeat) {
        case -1:
        case 1:
            this.repeat = 0;
            /* falls through */

        case 0:
            this.goto(this.q[this.q.length - 1].qid);
        }
    };

    this.next = function() {
        let song = this.q.shift();

        switch (this.repeat) {
        case -1:
            return;

        case 1:
            this.repeat = 0;
            /* falls through */

        case 0:
            // TODO: cover the shuffle
            this.q.push(song);
        }
    };

    this.goto = function(qid) {
        while (this.q[0].qid != qid)
            this.next();
    };

    this.moveon = function() {
        if (this.repeat === -1 || this.repeat === 0)
            this.next();
    }

}
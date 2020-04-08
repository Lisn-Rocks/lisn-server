export function SongQueue() {

    // q contains songs in the order they were added.
    this.q = [];

    // count is used to assign unique QIDs to the songs being added
    this.count = 0;

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
        song.qid = this.count++;
        this.q.push(song);
    };
    
    this.prev = function() {
        /* eslint-disable no-console */
        console.log("prev");
        /* eslint-enable no-console */
        switch (this.repeat) {
        case -1:
        case 1:
            this.repeat = 0;
            /* falls through */

        case 0:
            this.goto(this.get(this.q.length - 1).qid);
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
        while (this.get(0).qid != qid)
            this.next();
    };

    // moveon is used to automatically go to the next song in the Queue when the
    // current one is over. Obviously, it only ever plays the next song if
    // repeat mode is not set to 'repeat one song'.
    this.moveon = function() {
        if (this.repeat === -1 || this.repeat === 0)
            this.next();
    }

}
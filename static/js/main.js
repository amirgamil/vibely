const {
    Record,
    StoreOf,
    Component,
    ListOf,
} = window.Torus;

class App extends Component {
    init() {
        
        this.searched = false;
        this.loadSong = false;
        this.inputValue = "";
        this.path = "";
        this.songResults = [];
        this.song = "";


        this.handleSearch = this.handleSearch.bind(this);
        this.generateLyrics = this.generateLyrics.bind(this);
        
    }

    remove() {
        super.remove();
    }

    handleSearch(evt) {
        this.searched = true;
        this.inputValue = evt.target.value;
        if (this.inputValue === "") {
            this.songResults = []
        } else {
            fetch("/searchSongs"+this.inputValue)
            .then(res => res.json())
            .then(data => {
                const hits = data.response.sections[0].hits
                this.songResults = hits.map((element, index) => ({title: element.result.full_title, path: element.result.path}));
            }).catch(exception => {
                console.log("Exception parsing JSON data from backend " + exception)
            });
        }

        this.render();
    }

    async generateLyrics(title, path) {
        this.songResults = [];
        this.loadSong = true;
        this.inputValue = title
        this.path = path
        this.render()
        try {
            const res = await fetch("/scramble" + this.path.substring(1));
            const data = await res.json();
            this.song = "";
            var i = 0;
            data.data.forEach(element => {
                if (element.hidden) {
                    this.song += "<span class='hidden'>" + element.word + "</span> ";
                } else {
                    this.song += element.word + " ";
                }
            });
        } catch (exception) {
            console.log("exception generating scrambled lyrics! " + exception);
        }
        this.render();
        console.log(this.song);
        document.getElementById("lyrics").innerHTML = this.song;
    }

    compose() {
		return jdom
        `<main class="app">
            <header>
                <div class="column-div">
                    <h1 class="block">🎤 Vibely</h1>
                    <p>You ever just wanna vibe but you <span class="highlight">don't know the lyrics of a song</span>. Like you're going to a concert (before the pandemic) or 
                    it's karaoke or even if you're jamming by yourself. But you don't know the lyrics... </p>
                    <br/>
                    <p>Vibely is the <span class="highlight">fastest way</span> to learn lyrics and get you to vibe town as quickly as possible</p>
                </div>
            </header>
            <div class="column-div">
                <div class="row-div">
                    <div class = "input-div wrapper block">
                        <input value=${this.inputValue} type="text" placeholder="Enter the name of a song you want to learn" oninput=${this.handleSearch}/>
                        
                        ${this.searched ? jdom`
                        <div class="autoCompleteDropDown">
                            ${this.songResults.map((element, id) => {
                                return jdom`<div class="dropdownElement" onclick=${(evt) => this.generateLyrics(element.title, element.path)}>
                                       ${element.title}</div>`
                            })}
                        </div>
                        ` : null}
                    </div>
                    <button class = "block" 
                        onclick=${(evt) => this.song.length > 0 ? this.generateLyrics(this.inputValue, this.path) : alert("select a song first")}>
                        Scramble
                    </button>
                </div>
                ${this.song.length > 0 ? jdom`
                    <div id="lyrics" class="block wrapper">
                    </div>
                ` :  null}
            </div>
                
            <footer>
                <p>
                    Built with love by <a title="Visit the blog of Amir Bolous" href = "https://amirbolous.com">Amir</a> and is avaliable <a href = "https://github.com/amirgamil/vibely" title="View the GitHub repository for Vibely">open source</a>
                </p>
            </footer>
       </main>`;
    }
}

const app = new App();
document.body.appendChild(app.node);

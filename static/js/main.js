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
        this.songResults = [];


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


    generateLyrics(element) {
        this.songResults = [];
        this.loadSong = true;
        this.inputValue = element.title;
        fetch("/getSong" + element.path.substring(1))
        this.render()
    }

    compose() {
		return jdom
        `<main class="app">
            <header>
                <div class="center-div">
                    <h1 class="block">🎤 Vibely</h1>
                    <p>You ever just wanna vibe but you <span>don't know the lyrics of a song</span>. Like you're going to a concert (before the pandemic) or 
                    it's karaoke or even if you're jamming by yourself. But you don't know the lyrics... </p>
                    <br/>
                    <p>Vibely is the <span>fastest way</span> to learn lyrics and get you to vibe town as quickly as possible</p>
                </div>
            </header>
            <div class="center-div">
                <div class = "input-div wrapper block">
                    <input value=${this.inputValue} type="text" placeholder="Enter the name of a song you want to learn" oninput=${this.handleSearch}/>
                    ${this.searched ? jdom`
                    <div class="autoCompleteDropDown">
                        ${this.songResults.map((element, id) => {
                            return jdom`<div class="dropdownElement" onclick=${(evt) => this.generateLyrics(element)}>${element.title}</div>`
                        })}
                    </div>
            
                    ` : null}
                </div>
               
                
                
            </div>
            <footer>
                <p>
                    Built with love by
                    <a href = "https://amirbolous.com">
                        Amir
                    </a>
                </p>
            </footer>
       </main>`;
    }
}

const app = new App();
document.body.appendChild(app.node);

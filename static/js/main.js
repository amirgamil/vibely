const {
    Record,
    StoreOf,
    Component,
    ListOf,
} = window.Torus;

class App extends Component {
    init() {
        
        this.searched = false;
        this.inputValue = "";
        this.songResults = ["a", "b", "c"];


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
                console.log(data);
                this.songResults = data.items.map((element, index) => element.album.name);
            }).catch(exception => {
                console.log("Exception parsing JSON data from backend " + exception)
            });
        }

        this.render();
    }


    generateLyrics(evt) {
        this.songResults = [];
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
                </div>
            </header>
            <div class="center-div">
                <div class = "input-div wrapper block">
                    <input value=${this.inputValue} type="text" placeholder="Enter the name of a song you want to learn" oninput=${this.handleSearch}/>
                    ${this.searched ? jdom`
                    <div class="autoCompleteDropDown">
                        ${this.songResults.map((element, id) => {
                            return jdom`<div class="dropdownElement" onclick=${this.generateLyrics}>${element}</div>`
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

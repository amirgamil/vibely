const {
    Record,
    StoreOf,
    Component,
    ListOf,
} = window.Torus;

class App extends Component {
    init() {
		

    }

    remove() {
        super.remove();
    }

    compose() {
		return jdom
        `<main class="app">
            <header>
                <h1>🎤 Vibely</h1>
                <div class = "header-right">
                    <button class = "add" onclick=${() => {
                        this.store.create({h: '', b: '', t: []});
                        console.log(this.store.summarize());
                    }}>
                    +
                    </button>
                </div>

            </header>

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

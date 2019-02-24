import m from "mithril"

const root = document.body

m.render(root, 
    <div>
        <div id="sidebar-area">
            <div class="buttons">
                <a class="button is-danger">⏹</a>
                <a class="button is-light">▶</a>
            </div>
            <ul class="keywords">
                <li>foo</li>
                <li>foo</li>
                <li>foo</li>
            </ul>
        </div>
        <div id="main-area">
            <div id="top-bar">
                <input class="input is-rounded" type="text" placeholder="http://"></input>
            </div>
            <div id="iframe-browser">
                <iframe src=""></iframe>
            </div>
        </div>
    </div>
)

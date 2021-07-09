package main

func genlistall() string {
	html := ""
	for i := 0; i < 5; i++ {
		link := "/fav"
		icon := "http://localhost/img/icon.png"
		title := "タイトル"
		category := "カテゴリ"
		overview := "概要"
		stars := "星の数"
		html += "<a href=\"" + link + "\">	<li class=\"inline-block border-b border-gray-300 flex justify-between items-center py-4\">		<div class=\"flex items-start w-2/5\">			<div class=\"w-10 h-10 rounded mr-3\">				<img src=\"" + icon + "\" class=\"rounded-full h-10 w-10 bg-gray-300 m-auto\">			</div>			<div class=\"flex-1 overflow-hidden\">				<div>					<span class=\"font-bold \">" + title + "</span>				</div>				<p class=\"text-gray-700 text-xs\">" + category + "</p>			</div>		</div>		<p class=\"w-2/5\">" + overview + "</p>		<label for=\"stars\" class=\"font-bold w-1/5 text-right\">" + stars + "</label>	</li></a>\n"
	}
	return html
}

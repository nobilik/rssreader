package rssreader

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var fakeChannel = `<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title>
            <![CDATA[CNN.com - RSS Channel - HP Hero]]>
        </title>
        <description>
            <![CDATA[CNN.com delivers up-to-the-minute news and information on the latest top stories, weather, entertainment, politics and more.]]>
        </description>
        <link>https://www.cnn.com/index.html</link>
        <image>
            <url>http://i2.cdn.turner.com/cnn/2015/images/09/24/cnn.digital.png</url>
            <title>CNN.com - RSS Channel - HP Hero</title>
            <link>https://www.cnn.com/index.html</link>
        </image>
        <generator>coredev-bumblebee</generator>
        <lastBuildDate>Mon, 06 Nov 2023 13:08:05 GMT</lastBuildDate>
        <pubDate>Tue, 18 Apr 2023 21:25:59 GMT</pubDate>
        <copyright>
            <![CDATA[Copyright (c) 2023 Turner Broadcasting System, Inc. All Rights Reserved.]]>
        </copyright>
        <language>
            <![CDATA[en-US]]>
        </language>
        <ttl>10</ttl>
        <item>
            <title>
                <![CDATA[Some on-air claims about Dominion Voting Systems were false, Fox News acknowledges in statement after deal is announced]]>
            </title>
            <link>https://www.cnn.com/business/live-news/fox-news-dominion-trial-04-18-23/index.html</link>
			<description>
                <![CDATA[Fox-Dominion]]>
            </description>
            <guid isPermaLink="true">https://www.cnn.com/business/live-news/fox-news-dominion-trial-04-18-23/index.html</guid>
            <pubDate>Wed, 19 Apr 2023 12:44:51 GMT</pubDate>
            <media:group>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-super-169.jpg" height="619" width="1100" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-large-11.jpg" height="300" width="300" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-vertical-large-gallery.jpg" height="552" width="414" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-video-synd-2.jpg" height="480" width="640" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-live-video.jpg" height="324" width="576" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-vertical-gallery.jpg" height="360" width="270" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-story-body.jpg" height="169" width="300" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-assign.jpg" height="186" width="248" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230418164538-02-dominion-fox-trial-settlement-0418-hp-video.jpg" height="144" width="256" type="image/jpeg"/>
            </media:group>
        </item>
        <item>
            <title>
                <![CDATA[Dominion still has pending lawsuits against election deniers such as Rudy Giuliani and Sidney Powell]]>
            </title>
            <link>https://www.cnn.com/business/live-news/fox-news-dominion-trial-04-18-23/h_8d51e3ae2714edaa0dace837305d03b8</link>
            <guid isPermaLink="true">https://www.cnn.com/business/live-news/fox-news-dominion-trial-04-18-23/h_8d51e3ae2714edaa0dace837305d03b8</guid>
            <media:group>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-super-169.jpg" height="619" width="1100" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-large-11.jpg" height="300" width="300" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-vertical-large-gallery.jpg" height="552" width="414" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-video-synd-2.jpg" height="480" width="640" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-live-video.jpg" height="324" width="576" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-vertical-gallery.jpg" height="360" width="270" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-story-body.jpg" height="169" width="300" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-assign.jpg" height="186" width="248" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230417170417-fox-news-headquarters-0228-hp-video.jpg" height="144" width="256" type="image/jpeg"/>
            </media:group>
        </item>
        <item>
            <title>
                <![CDATA[Here are the 20 specific Fox broadcasts and tweets Dominion says were defamatory]]>
            </title>
            <description>
                <![CDATA[• Fox-Dominion trial delay 'is not unusual,' judge says
• Fox News' defamation battle isn't stopping Trump's election lies]]>
            </description>
            <link>https://www.cnn.com/2023/04/17/media/dominion-fox-news-allegations/index.html</link>
            <guid isPermaLink="true">https://www.cnn.com/2023/04/17/media/dominion-fox-news-allegations/index.html</guid>
            <pubDate>Mon, 17 Apr 2023 16:01:11 GMT</pubDate>
            <media:group>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-super-169.jpg" height="619" width="1100" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-large-11.jpg" height="300" width="300" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-vertical-large-gallery.jpg" height="552" width="414" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-video-synd-2.jpg" height="480" width="640" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-live-video.jpg" height="324" width="576" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-vertical-gallery.jpg" height="360" width="270" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-story-body.jpg" height="169" width="300" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-t1-main.jpg" height="250" width="250" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-assign.jpg" height="186" width="248" type="image/jpeg"/>
                <media:content medium="image" url="https://cdn.cnn.com/cnnnext/dam/assets/230416140110-maria-bartiromo-file-020823-restricted-hp-video.jpg" height="144" width="256" type="image/jpeg"/>
            </media:group>
        </item>
    </channel>
</rss>`

var fakeItems = `
<?xml version="1.0" encoding="UTF-8"?>
<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:prism="http://prismstandard.org/namespaces/basic/2.0/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:syn="http://purl.org/rss/1.0/modules/syndication/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns="http://purl.org/rss/1.0/">
    <channel rdf:about="http://www.chinaminingmagazine.com/">
        <title>中国矿业</title>
        <link>http://www.chinaminingmagazine.com/</link>
        <description>CHINA MINING MAGAZINE</description>
        <syn:updatePeriod>day</syn:updatePeriod>
        <syn:updateFrequency>15</syn:updateFrequency>
        <syn:updateBase>2023-11-03 17:11:14</syn:updateBase>
        <dc:creator>
            <a href="mailto:zgkyzzs@163.com">zgkyzzs@163.com</a>
        </dc:creator>
        <dc:publisher>
            <a href="mailto:zgkyzzs@163.com">zgkyzzs@163.com</a>
        </dc:publisher>
        <dc:date>2023-11-03 17:11:14</dc:date>
        <dc:language>zh</dc:language>
        <dc:rights>《中国矿业》杂志社有限公司</dc:rights>
        <prism:copyright>《中国矿业》杂志社有限公司</prism:copyright>
        <prism:rightsAgent>
            <a href="mailto:zgkyzzs@163.com">zgkyzzs@163.com</a>
        </prism:rightsAgent>
        <prism:issn>1004-4051</prism:issn>
        <items>
            <rdf:Seq>
                <rdf:li rdf:resource="http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230662"/>
                <rdf:li rdf:resource="http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230658"/>
                <rdf:li rdf:resource="http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230677"/>
            </rdf:Seq>
        </items>
    </channel>
    <item rdf:about="http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230662">
        <title>
            <![CDATA[全球铜矿资源分布现状及勘查投入分析]]>
        </title>
        <link>http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230662</link>
        <description>
            <![CDATA[
        	                	
                			&lt;br/&gt;&lt;p&gt;&amp;lt;p&amp;lt;铜是一种不可再生的金属矿产资源，2016年被自然资源部（原国土资源部）列入战略性矿产目录，是我国对外依存度较高的金属矿产之一。本文介绍了全球铜矿资源分布、铜矿床类型、大型铜矿公司、铜矿勘查投入等情况。全球的铜矿资源主要分布在智利、澳大利亚、秘鲁等国；铜矿床类型主要为斑岩型；大型铜矿公司铜产量占全球铜产量的比重较高；近十年来全球铜勘查投入趋势呈“W”型，拉丁美洲是全球铜勘查投入最高的地区，矿山阶段的铜勘查投入近十年来，也是有记录以来第一次超过了其他任何阶段的预算，未来数年矿山阶段的铜勘查投入占比可能仍然会保持较高的位置。&amp;lt;/p&amp;lt;&lt;/p&gt;
			&lt;br/&gt;中国矿业. 2023 32(S2): 1-6. 		]]>
	
        </description>
        <content:encoded>
            <![CDATA[
            	                	
                			&lt;br/&gt;&lt;p&gt;&amp;lt;p&amp;lt;铜是一种不可再生的金属矿产资源，2016年被自然资源部（原国土资源部）列入战略性矿产目录，是我国对外依存度较高的金属矿产之一。本文介绍了全球铜矿资源分布、铜矿床类型、大型铜矿公司、铜矿勘查投入等情况。全球的铜矿资源主要分布在智利、澳大利亚、秘鲁等国；铜矿床类型主要为斑岩型；大型铜矿公司铜产量占全球铜产量的比重较高；近十年来全球铜勘查投入趋势呈“W”型，拉丁美洲是全球铜勘查投入最高的地区，矿山阶段的铜勘查投入近十年来，也是有记录以来第一次超过了其他任何阶段的预算，未来数年矿山阶段的铜勘查投入占比可能仍然会保持较高的位置。&amp;lt;/p&amp;lt;&lt;/p&gt;
			&lt;br/&gt;中国矿业. 2023 32(S2): 1-6. 		]]>
	
        </content:encoded>
        <dc:title>
            <![CDATA[全球铜矿资源分布现状及勘查投入分析]]>
        </dc:title>
        <dc:creator>
            <![CDATA[                	
                ]]>
        </dc:creator>
        <dc:date></dc:date>
        <dc:rights>Personal use only, all commercial or other reuse prohibited</dc:rights>
        <dc:source>. 2023 32(S2): 1-6.</dc:source>
        <dc:type>article</dc:type>
        <dc:identifier>doi:10.12075/j.issn.1004-4051.20230662</dc:identifier>
        <prism:doi>10.12075/j.issn.1004-4051.20230662</prism:doi>
        <prism:publicationName>中国矿业</prism:publicationName>
        <prism:volume>32</prism:volume>
        <prism:number>S2</prism:number>
        <prism:publicationDate></prism:publicationDate>
        <prism:url>http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230662</prism:url>
        <prism:startingPage>1-6</prism:startingPage>
    </item>
    <item rdf:about="http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230658">
        <title>
            <![CDATA[湘黔桂三角区层控型铅锌矿成矿规律和找矿标志]]>
        </title>
        <link>http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230658</link>
        <description>
            <![CDATA[
        	                	
                			&lt;br/&gt;&lt;p&gt;&amp;lt;p&amp;lt;湘黔桂三角区层控型铅锌矿受深大断裂和赋矿岩性层两条主线控制，隆起西侧松桃-三都深断裂带控制铅锌成矿和铅锌矿带的空间展布，形成长大于500 km的铅锌汞多金属成矿带。矿化多分布在区域深大断裂夹持地带和深大断裂两侧5～7 km范围狭长地带，铅锌矿层位、岩性控制明显，层控型铅锌控矿层位主要为寒武系、次奥陶系、震旦系。此外，还对湘黔桂地区的成矿规律及找矿标志进行了阐述和总结。&amp;lt;/p&amp;lt;&lt;/p&gt;
			&lt;br/&gt;中国矿业. 2023 32(S2): 101-105. 		]]>
	
        </description>
        <content:encoded>
            <![CDATA[
            	                	
                			&lt;br/&gt;&lt;p&gt;&amp;lt;p&amp;lt;湘黔桂三角区层控型铅锌矿受深大断裂和赋矿岩性层两条主线控制，隆起西侧松桃-三都深断裂带控制铅锌成矿和铅锌矿带的空间展布，形成长大于500 km的铅锌汞多金属成矿带。矿化多分布在区域深大断裂夹持地带和深大断裂两侧5～7 km范围狭长地带，铅锌矿层位、岩性控制明显，层控型铅锌控矿层位主要为寒武系、次奥陶系、震旦系。此外，还对湘黔桂地区的成矿规律及找矿标志进行了阐述和总结。&amp;lt;/p&amp;lt;&lt;/p&gt;
			&lt;br/&gt;中国矿业. 2023 32(S2): 101-105. 		]]>
	
        </content:encoded>
        <dc:title>
            <![CDATA[湘黔桂三角区层控型铅锌矿成矿规律和找矿标志]]>
        </dc:title>
        <dc:creator>
            <![CDATA[                	
                ]]>
        </dc:creator>
        <dc:date></dc:date>
        <dc:rights>Personal use only, all commercial or other reuse prohibited</dc:rights>
        <dc:source>. 2023 32(S2): 101-105.</dc:source>
        <dc:type>article</dc:type>
        <dc:identifier>doi:10.12075/j.issn.1004-4051.20230658</dc:identifier>
        <prism:doi>10.12075/j.issn.1004-4051.20230658</prism:doi>
        <prism:publicationName>中国矿业</prism:publicationName>
        <prism:volume>32</prism:volume>
        <prism:number>S2</prism:number>
        <prism:publicationDate></prism:publicationDate>
        <prism:url>http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230658</prism:url>
        <prism:startingPage>101-105</prism:startingPage>
    </item>
    <item rdf:about="http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230677">
        <title>
            <![CDATA[西藏象背山铜多金属矿地质特征及找矿潜力分析]]>
        </title>
        <link>http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230677</link>
        <description>
            <![CDATA[
        	                	
                			&lt;br/&gt;&lt;p&gt;&amp;lt;p&amp;lt;西藏象背山铜多金属矿位于冈底斯斑岩-矽卡岩铜多金属成矿带东段，夹持于甲玛和驱龙两个超大型斑岩-矽卡岩矿床之间。象背山铜多金属矿为斑岩-矽卡岩矿床，是甲玛多中心复合斑岩成矿系统的一部分，其矿化类型主要为矽卡岩矿化和斑岩矿化，主成矿元素为铜、钼，伴生金、银；热液蚀变主要矽卡岩化、硅化、绢英岩化。其中，矽卡岩矿体主要赋存于斑岩和大理岩接触带以及角岩和大理岩接触带。成矿地质条件研究、地球化学勘查和高光谱蚀变提取成果指示，象背山铜多金属矿具有形成大型斑岩-矽卡岩铜钼（金、银）的潜力。&amp;lt;/p&amp;lt;&lt;/p&gt;
			&lt;br/&gt;中国矿业. 2023 32(S2): 106-113. 		]]>
	
        </description>
        <content:encoded>
            <![CDATA[
            	                	
                			&lt;br/&gt;&lt;p&gt;&amp;lt;p&amp;lt;西藏象背山铜多金属矿位于冈底斯斑岩-矽卡岩铜多金属成矿带东段，夹持于甲玛和驱龙两个超大型斑岩-矽卡岩矿床之间。象背山铜多金属矿为斑岩-矽卡岩矿床，是甲玛多中心复合斑岩成矿系统的一部分，其矿化类型主要为矽卡岩矿化和斑岩矿化，主成矿元素为铜、钼，伴生金、银；热液蚀变主要矽卡岩化、硅化、绢英岩化。其中，矽卡岩矿体主要赋存于斑岩和大理岩接触带以及角岩和大理岩接触带。成矿地质条件研究、地球化学勘查和高光谱蚀变提取成果指示，象背山铜多金属矿具有形成大型斑岩-矽卡岩铜钼（金、银）的潜力。&amp;lt;/p&amp;lt;&lt;/p&gt;
			&lt;br/&gt;中国矿业. 2023 32(S2): 106-113. 		]]>
	
        </content:encoded>
        <dc:title>
            <![CDATA[西藏象背山铜多金属矿地质特征及找矿潜力分析]]>
        </dc:title>
        <dc:creator>
            <![CDATA[                	
                ]]>
        </dc:creator>
        <dc:date></dc:date>
        <dc:rights>Personal use only, all commercial or other reuse prohibited</dc:rights>
        <dc:source>. 2023 32(S2): 106-113.</dc:source>
        <dc:type>article</dc:type>
        <dc:identifier>doi:10.12075/j.issn.1004-4051.20230677</dc:identifier>
        <prism:doi>10.12075/j.issn.1004-4051.20230677</prism:doi>
        <prism:publicationName>中国矿业</prism:publicationName>
        <prism:volume>32</prism:volume>
        <prism:number>S2</prism:number>
        <prism:publicationDate></prism:publicationDate>
        <prism:url>http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230677</prism:url>
        <prism:startingPage>106-113</prism:startingPage>
    </item>
</rdf:RDF>
`

// go test -race
func TestParse(t *testing.T) {
	urlsCount := 10
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakeChannel))
	}))
	defer server.Close()

	urls := []string{}

	for i := 0; i < urlsCount; i++ {
		urls = append(urls, server.URL)
	}

	rssItems, errors := Parse(urls...)
	if len(errors) > 0 {
		t.Errorf("Error(s) occurred while parsing the RSS feed: %v", errors)
	}

	if len(rssItems) != urlsCount*3 {
		t.Errorf("Expected %d RSS item, got %d", urlsCount*3, len(rssItems))
	}

	expectedItem := RssItem{
		Title:       "Some on-air claims about Dominion Voting Systems were false, Fox News acknowledges in statement after deal is announced",
		Source:      "CNN.com - RSS Channel - HP Hero",
		SourceURL:   "https://www.cnn.com/index.html",
		Link:        "https://www.cnn.com/business/live-news/fox-news-dominion-trial-04-18-23/index.html",
		Description: "Fox-Dominion",
	}
	expectedItem.PublishDate, _ = parseTime("Wed, 19 Apr 2023 12:44:51 GMT")

	if rssItems[0].Title != expectedItem.Title {
		t.Errorf("Title mismatch. Expected: %s, Actual: %s", expectedItem.Title, rssItems[0].Title)
	}

	if rssItems[0].Source != expectedItem.Source {
		t.Errorf("Source mismatch. Expected: %s, Actual: %s", expectedItem.Source, rssItems[0].Source)
	}

	if rssItems[0].SourceURL != expectedItem.SourceURL {
		t.Errorf("SourceURL mismatch. Expected: %s, Actual: %s", expectedItem.SourceURL, rssItems[0].SourceURL)
	}

	if rssItems[0].Link != expectedItem.Link {
		t.Errorf("Link mismatch. Expected: %s, Actual: %s", expectedItem.Link, rssItems[0].Link)
	}

	if !rssItems[0].PublishDate.Equal(expectedItem.PublishDate) {
		t.Errorf("PublishDate mismatch. Expected: %s, Actual: %s", expectedItem.PublishDate, rssItems[0].PublishDate)
	}

	if rssItems[0].Description != expectedItem.Description {
		t.Errorf("Description mismatch. Expected: %s, Actual: %s", expectedItem.Description, rssItems[0].Description)
	}
}

func TestParseWithItems(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakeItems))
	}))
	defer server.Close()

	rssItems, errors := Parse(server.URL)
	if len(errors) > 0 {
		t.Errorf("Error(s) occurred while parsing the RSS feed: %v", errors)
	}

	if len(rssItems) != 3 {
		t.Errorf("Expected 3 RSS item, got %d", len(rssItems))
	}

	expectedItem := RssItem{
		Title:     "全球铜矿资源分布现状及勘查投入分析",
		Source:    "中国矿业",
		SourceURL: "http://www.chinaminingmagazine.com/",
		Link:      "http://www.chinaminingmagazine.com/article/doi/10.12075/j.issn.1004-4051.20230662",
		Description: `&lt;br/&gt;&lt;p&gt;&amp;lt;p&amp;lt;铜是一种不可再生的金属矿产资源，2016年被自然资源部（原国土资源部）列入战略性矿产目录，是我国对外依存度较高的金属矿产之一。本文介绍了全球铜矿资源分布、铜矿床类型、大型铜矿公司、铜矿勘查投入等情况。全球的铜矿资源主要分布在智利、澳大利亚、秘鲁等国；铜矿床类型主要为斑岩型；大型铜矿公司铜产量占全球铜产量的比重较高；近十年来全球铜勘查投入趋势呈“W”型，拉丁美洲是全球铜勘查投入最高的地区，矿山阶段的铜勘查投入近十年来，也是有记录以来第一次超过了其他任何阶段的预算，未来数年矿山阶段的铜勘查投入占比可能仍然会保持较高的位置。&amp;lt;/p&amp;lt;&lt;/p&gt;
			&lt;br/&gt;中国矿业. 2023 32(S2): 1-6.`,
	}
	expectedItem.PublishDate, _ = parseTime("0001-01-01 00:00:00 +0000 UTC")

	if rssItems[0].Title != expectedItem.Title {
		t.Errorf("Title mismatch. Expected: %s, Actual: %s", expectedItem.Title, rssItems[0].Title)
	}

	if rssItems[0].Source != expectedItem.Source {
		t.Errorf("Source mismatch. Expected: %s, Actual: %s", expectedItem.Source, rssItems[0].Source)
	}

	if rssItems[0].SourceURL != expectedItem.SourceURL {
		t.Errorf("SourceURL mismatch. Expected: %s, Actual: %s", expectedItem.SourceURL, rssItems[0].SourceURL)
	}

	if rssItems[0].Link != expectedItem.Link {
		t.Errorf("Link mismatch. Expected: %s, Actual: %s", expectedItem.Link, rssItems[0].Link)
	}

	if !rssItems[0].PublishDate.Equal(expectedItem.PublishDate) {
		t.Errorf("PublishDate mismatch. Expected: %s, Actual: %s", expectedItem.PublishDate, rssItems[0].PublishDate)
	}

	if rssItems[0].Description != expectedItem.Description {
		t.Errorf("Description mismatch. Expected: %s, Actual: %s", expectedItem.Description, rssItems[0].Description)
	}
}

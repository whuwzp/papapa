package paser

const CityString = `<a href="(http://city.zhenai.com/[a-z0-9]+)"[^>]*>([^<]+).*`
const MenString = `<td><a href="(http://album.zhenai.com/u/[a-z0-9]+)"[^>]*>([^<]+).*`
const ManString = ``

const NextMen = `<a class="next-page"[^h]*href="(http://city.zhenai.com/zhumadian/[0-9]+)"[^>]*>下一页<`

//<a class="next-page" href="http://city.zhenai.com/zhumadian/3">下一页</a>
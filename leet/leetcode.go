package leet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gladmo/leetcode/store"
)

const UA = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`

func getCookie() (cookie []*http.Cookie, err error) {
	URL := "https://leetcode-cn.com/graphql/"
	payload := strings.NewReader(`{"operationName":"globalData","variables":{},"query":"query globalData {\n  feature {\n    questionTranslation\n    subscription\n    signUp\n    discuss\n    mockInterview\n    contest\n    store\n    book\n    chinaProblemDiscuss\n    socialProviders\n    studentFooter\n    cnJobs\n    enableLsp\n    enableWs\n    enableDebugger\n    enableDebuggerAdmin\n    enableDarkMode\n    tasks\n    leetbook\n    enableEduDiscount\n    __typename\n  }\n  userStatus {\n    isSignedIn\n    isAdmin\n    isStaff\n    isSuperuser\n    isTranslator\n    isPremium\n    isVerified\n    isPhoneVerified\n    isWechatVerified\n    checkedInToday\n    username\n    realName\n    userSlug\n    groups\n    avatar\n    optedIn\n    requestRegion\n    region\n    socketToken\n    activeSessionId\n    permissions\n    notificationStatus {\n      lastModified\n      numUnread\n      __typename\n    }\n    completedFeatureGuides\n    useTranslation\n    accountStatus {\n      isFrozen\n      inactiveAfter\n      __typename\n    }\n    __typename\n  }\n  siteRegion\n  chinaHost\n  websocketUrl\n  userBannedInfo {\n    bannedData {\n      endAt\n      bannedType\n      __typename\n    }\n    __typename\n  }\n  commonNojPermissionTypes\n  jobsMyCompany {\n    name\n    __typename\n  }\n  myLastCompanyClaimApplication(status: PENDING) {\n    id\n    name\n    countryCode\n    phone\n    department\n    jobTitle\n    createdAt\n    opinion\n    reviewedAt\n    claimType\n    status\n    companyTag {\n      name\n      translatedName\n      slug\n      __typename\n    }\n    authorization\n    businessLicense\n    __typename\n  }\n}\n"}`)

	req, err := http.NewRequest(http.MethodPost, URL, payload)
	if err != nil {
		return
	}
	req.Header.Add("authority", "leetcode-cn.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("x-timezone", "Asia/Shanghai")
	req.Header.Add("x-operation-name", "globalData")
	req.Header.Add("accept-language", "zh-CN")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "*/*")
	req.Header.Add("x-csrftoken", "undefined")
	req.Header.Add("x-definition-name", "feature,userStatus,siteRegion,chinaHost,websocketUrl,userBannedInfo,commonNojPermissionTypes,jobsMyCompany,myLastCompanyClaimApplication")
	req.Header.Add("origin", "https://leetcode-cn.com")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://leetcode-cn.com/")

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	cookie = res.Cookies()
	return
}

const (
	cacheDuration = time.Hour * 24
)

// Parse ...
// param: 222
// param: https://leetcode-cn.com/problems/k-th-symbol-in-grammar
// param: https://leetcode-cn.com/problems/k-th-symbol-in-grammar/solution/
// param: leetcode-cn.com/problems/k-th-symbol-in-grammar
func Parse(param string) string {
	if match, err := regexp.MatchString(`\d+`, param); err == nil && match {
		stat, err := ProblemID2name(param)
		if err == nil {
			return stat.QuestionTitleSlug
		}
	}

	return ParseFromURL(param)
}

// ParseFromURL
// param: https://leetcode-cn.com/problems/k-th-symbol-in-grammar
// param: https://leetcode-cn.com/problems/k-th-symbol-in-grammar/solution/
// param: leetcode-cn.com/problems/k-th-symbol-in-grammar
func ParseFromURL(param string) string {
	if strings.HasPrefix(param, "http") ||
		strings.HasPrefix(param, "leetcode-cn.com") {

		re := regexp.MustCompile(`leetcode-cn\.com/problems/(.*)`)

		result := re.FindStringSubmatch(param)
		if len(result) == 2 {
			result := strings.Trim(result[1], "/")
			ss := strings.Split(result, "/")
			if len(ss) == 2 {
				result = ss[0]
			}
			return result
		}
	}

	return ""
}

// ProblemID2name return problem info
func ProblemID2name(id string) (stats store.QuestionStats, err error) {
	if ok, err := store.ProblemInfoIsExpire(); err == nil && !ok {
		stats, err = store.GetProblemsInfo(id)
		if err == nil {
			return stats, err
		}
	}

	allProblemsResult, err := allProblems()
	if err != nil {
		return
	}

	var infos []store.QuestionStats
	for _, pair := range allProblemsResult.StatStatusPairs {
		infos = append(infos, pair.Stat)
	}

	err = store.UpdateProblems(infos)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = store.ProblemsTTL(time.Now().Add(cacheDuration))
	if err != nil {
		return
	}

	return store.GetProblemsInfo(id)
}

func allProblems() (apr AllProblemsResult, err error) {
	URL := "https://leetcode-cn.com/api/problems/all/"

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return
	}
	req.Header.Add("authority", "leetcode-cn.com")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("user-agent", UA)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://leetcode-cn.com/problemset/all/")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &apr)
	if err != nil {
		return
	}

	return
}

type AllProblemsResult struct {
	UserName        string `json:"user_name"`
	NumSolved       int    `json:"num_solved"`
	NumTotal        int    `json:"num_total"`
	AcEasy          int    `json:"ac_easy"`
	AcMedium        int    `json:"ac_medium"`
	AcHard          int    `json:"ac_hard"`
	StatStatusPairs []struct {
		Stat       store.QuestionStats `json:"stat"`
		Status     interface{}         `json:"status"`
		Difficulty struct {
			Level int `json:"level"`
		} `json:"difficulty"`
		PaidOnly  bool `json:"paid_only"`
		IsFavor   bool `json:"is_favor"`
		Frequency int  `json:"frequency"`
		Progress  int  `json:"progress"`
	} `json:"stat_status_pairs"`
	FrequencyHigh int    `json:"frequency_high"`
	FrequencyMid  int    `json:"frequency_mid"`
	CategorySlug  string `json:"category_slug"`
}

func Fetch(title string) (qd QuestionDetail, err error) {
	URL := "https://leetcode-cn.com/graphql/"

	payload := strings.NewReader(fmt.Sprintf(`{
    "operationName": "questionData",
    "variables": {
        "titleSlug": "%s"
    },
    "query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    envInfo\n    book {\n      id\n      bookName\n      pressName\n      source\n      shortDescription\n      fullDescription\n      bookImgUrl\n      pressImgUrl\n      productUrl\n      __typename\n    }\n    isSubscribed\n    isDailyQuestion\n    dailyRecordStatus\n    editorType\n    ugcQuestionId\n    style\n    __typename\n  }\n}\n"
}`, title))
	req, err := http.NewRequest(http.MethodPost, URL, payload)
	if err != nil {
		return
	}

	cookies, err := getCookie()
	if err != nil {
		return
	}

	var csrfToken string
	for _, cookie := range cookies {
		if cookie.Name == "csrftoken" {
			csrfToken = cookie.Value
		}

		req.AddCookie(cookie)
	}

	req.Header.Add("authority", "leetcode-cn.com")
	req.Header.Add("x-timezone", "Asia/Shanghai")
	req.Header.Add("x-operation-name", "questionData")
	req.Header.Add("accept-language", "zh-CN")
	req.Header.Add("user-agent", UA)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "*/*")
	req.Header.Add("x-csrftoken", csrfToken)
	req.Header.Add("x-definition-name", "question")
	req.Header.Add("origin", "https://leetcode-cn.com")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://leetcode-cn.com")

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &qd)
	if err != nil {
		return
	}

	return
}

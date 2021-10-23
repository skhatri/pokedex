#include <iostream>
#include <curl/curl.h>
#include "json.hpp"

using json = nlohmann::json;

using namespace std;

namespace poke {

    struct Data {
        std::string name;
        int s_no;
    };

    void to_json(json& j, const Data& d) {
        j = json{{"s_no", d.s_no}, {"name", d.name}};
    }

    void from_json(const json& j, Data& d) {
        j.at("s_no").get_to(d.s_no);
        j.at("name").get_to(d.name);
    }

    struct Metadata {
        int total;
    };

    void to_json(json& j, const Metadata& meta) {
        j = json {{"total", meta.total}};
    }

    void from_json(const json& j, Metadata& meta) {
        j.at("total").get_to(meta.total);
    }

    struct Result {
        Metadata metadata;
        std::vector<Data> data;
    };

    void to_json(json& j, const Result& res) {
        j = json {{"metadata", res.metadata}, {"data", res.data}};
    }
    void from_json(const json& j, Result& res) {
        j.at("metadata").get_to(res.metadata);
        const json& d = j.at("data");
        res.data.resize(d.size());
        std::copy(d.begin(), d.end(), res.data.begin());
    }
}

struct httpclient {

    size_t receive_data(char *buffer, size_t itemsize, size_t n, void *ignore){
        size_t len = itemsize * n;
        for (int i=0; i<len; i++) {
            printf("%c", buffer[i]);
        }
        json j = json::parse(buffer);
        auto res = j.get<poke::Result>();
        printf("total pokemons %d\n", res.metadata.total);
        cout << "name " << res.data.at(0).name << endl;
        return len;
    }
};

size_t receive_data(char *buffer, size_t itemsize, size_t n, void *ignore) {
    httpclient* h {};
    return h->receive_data(buffer, itemsize, n, ignore);
}

int main() {
    CURL *curl = curl_easy_init();

    if (!curl) {
        fprintf(stderr, "init curl");
        return EXIT_FAILURE;
    }
    curl_easy_setopt(curl, CURLOPT_URL, "http://localhost:6100/search/by-name?q=Dragapult");


    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, receive_data);

    CURLcode result = curl_easy_perform(curl);
    if (result != CURLE_OK) {
        fprintf(stderr, "issue accessing url");
    }

    curl_easy_cleanup(curl);

    return 0;
}

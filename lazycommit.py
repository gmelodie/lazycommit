import http.client
from git import Repo

def get_commit_msg():
    conn = http.client.HTTPConnection("whatthecommit.com")
    conn.request("GET", "/")
    response = conn.getresponse()

    if response.status != 200:
        print(f"Unable to get page, reason: {response.reason}")
        exit(1)


    html = response.read().decode()
    commit_msg = html.split("<p>")[1].split("</p>")[0].strip()

    return commit_msg


if __name__ == "__main__":
    repo = Repo(".")
    repo.git.add(all=True)
    print("git add --all")
    commit_msg = get_commit_msg()
    repo.index.commit(commit_msg)
    print(f"git commit -m \"{commit_msg}\"")
    info = repo.remotes.origin.push()[0]
    print(info.summary)

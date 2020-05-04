#!/bin/bash

#!/bin/bash

# Run this after using go-bin-generic as a template

remotes ()
{
    for x in $(git remote);
    do
        echo -n "$x ";
        git remote get-url $x;
    done
}

if [[ $1 ]]; then
    ORG=$1
    if [[ $2 ]]; then
        REPO=$2
    else
        echo "Checking remote for Repo"
        REPO=$(remotes | grep origin | awk -F '/' '{print $5}' | awk -F '.git' '{print $1}')
    fi
else
    echo "No args passed, checking remotes"
    ORG=$(remotes | grep origin | awk -F '/' '{print $4}' | awk -F '.git' '{print $1}')
    REPO=$(remotes | grep origin | awk -F '/' '{print $5}' | awk -F '.git' '{print $1}')
fi

printf "\n\tORG:\t${ORG}\n"
printf "\tREPO:\t${REPO}\n\n"

oldOrg=Benbentwo
oldRepo=go-bin-generic

run_arch=$(uname)
if [[ "${run_arch}" == "Darwin" ]]; then
    OS=mac
elif [[ "${run_arch}" == "Linux" ]]; then
    OS=linux
fi

#Begin Replacing
if [[ "${OS}" == "mac" ]]; then
    sed -i '' "s/${oldOrg}/${ORG}/" "go.mod"
else
    sed -i "s/${oldOrg}/${ORG}/1" "go.mod"
fi

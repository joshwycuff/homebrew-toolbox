class Typescore < Formula
  desc "Simple tool to score the typing difficulty of text"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/typescore-ed86cdab9fdea6be52d385c8a6eec4f0e050bf1d/typescore.tar.gz"
  version "ed86cdab9fdea6be52d385c8a6eec4f0e050bf1d"

  def install
    bin.install "typescore"
    man1.install Dir["*.1"]
    bash_completion.install "typescore.bash" => "typescore"
    zsh_completion.install "typescore.zsh" => "_typescore"
    fish_completion.install "typescore.fish"
  end
end
